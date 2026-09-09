use chrono::{DateTime, Utc};
use serde::{Deserialize, Deserializer, Serialize};
use url::Url;
use uuid::Uuid;

use domain::{
    Id,
    plugin::{PluginDraft, PluginFrontend, PluginName, PluginSlug, PluginView, ServiceEndpoint},
};

use crate::service::{Malformed, ServiceError, plugin_service::PluginChange};

use super::role::parse_permissions;

/// Represents an installed plugin, mirroring `PluginView` flat: no hash, no
/// key. A later task reads `frontend_target` directly off this response.
#[derive(Debug, Serialize, utoipa::ToSchema)]
#[schema(example = json!({
    "id": "01990000-0000-7000-8000-000000000001",
    "slug": "tbz-baumkataster",
    "name": "TBZ Baumkataster",
    "description": null,
    "organization_id": "01980000-0000-7000-8000-000000000001",
    "permissions": ["tree:create", "tree:update"],
    "required_permissions": ["tree:read"],
    "frontend_mode": "none",
    "frontend_target": null,
    "enabled": true,
    "has_credential": true,
    "last_seen_at": "2026-09-01T10:00:00Z",
    "created_at": "2026-08-01T00:00:00Z"
}))]
pub struct PluginResponse {
    pub id: Uuid,
    pub slug: String,
    pub name: String,
    pub description: Option<String>,
    pub organization_id: Uuid,
    pub permissions: Vec<String>,
    pub required_permissions: Vec<String>,
    pub frontend_mode: String,
    pub frontend_target: Option<String>,
    pub enabled: bool,
    pub has_credential: bool,
    pub last_seen_at: Option<DateTime<Utc>>,
    pub created_at: Option<DateTime<Utc>>,
}

impl From<&PluginView> for PluginResponse {
    fn from(view: &PluginView) -> Self {
        Self {
            id: view.id.value(),
            slug: view.slug.clone(),
            name: view.name.clone(),
            description: view.description.clone(),
            organization_id: view.organization_id.value(),
            permissions: view.permissions.clone(),
            required_permissions: view.required_permissions.clone(),
            frontend_mode: view.frontend_mode.to_string(),
            frontend_target: view.frontend_target.clone(),
            enabled: view.enabled,
            has_credential: view.has_credential,
            last_seen_at: view.last_seen_at,
            created_at: view.created_at,
        }
    }
}

/// One-time response carrying the plaintext API key, issued on install and
/// on rotation only — it never appears again afterwards.
#[derive(Debug, Serialize, utoipa::ToSchema)]
#[schema(example = json!({ "key": "gep_2f9a5c3e1b7d4a80.7b1e3c9a4f2d6081b3c5e7a9d1f3b5c7" }))]
pub struct PluginKeyResponse {
    pub key: String,
}

/// Frontend wiring for a plugin, tagged so mode and target cannot be set
/// inconsistently.
#[derive(Debug, Serialize, Deserialize, utoipa::ToSchema)]
#[serde(tag = "mode", rename_all = "lowercase")]
pub enum PluginFrontendDto {
    None,
    External { target: String },
    Proxied { target: String },
}

impl PluginFrontendDto {
    fn into_domain(self) -> Result<PluginFrontend, ServiceError> {
        match self {
            Self::None => Ok(PluginFrontend::None),
            Self::External { target } => {
                let url = Url::parse(&target).map_err(|e| ServiceError::Malformed {
                    kind: Malformed::PluginFrontend,
                    detail: format!("invalid url: {e}"),
                })?;
                // localhost is the one exception, for local development against
                // a plugin that has no certificate yet.
                let is_localhost = url.host_str() == Some("localhost");
                if url.scheme() != "https" && !is_localhost {
                    return Err(ServiceError::Malformed {
                        kind: Malformed::PluginFrontend,
                        detail:
                            "external frontend target must be an absolute https url (localhost excepted)"
                                .into(),
                    });
                }
                Ok(PluginFrontend::External(url))
            }
            Self::Proxied { target } => {
                let (host, port) =
                    target
                        .rsplit_once(':')
                        .ok_or_else(|| ServiceError::Malformed {
                            kind: Malformed::PluginFrontend,
                            detail: "proxied frontend target must be host:port".into(),
                        })?;
                let port: u16 = port.parse().map_err(|_| ServiceError::Malformed {
                    kind: Malformed::PluginFrontend,
                    detail: "proxied frontend target must carry a numeric port".into(),
                })?;
                // The service allowlist check arrives with the proxy plan; not
                // required here yet.
                Ok(PluginFrontend::Proxied(ServiceEndpoint::new(host, port)?))
            }
        }
    }
}

/// Request body for installing a new plugin.
#[derive(Debug, Deserialize, utoipa::ToSchema)]
#[schema(example = json!({
    "slug": "tbz-baumkataster",
    "name": "TBZ Baumkataster",
    "organization_id": "01980000-0000-7000-8000-000000000001",
    "permissions": ["tree:create"],
    "required_permissions": ["tree:read"],
    "frontend": { "mode": "none" }
}))]
pub struct PluginCreateRequest {
    pub slug: String,
    pub name: String,
    pub description: Option<String>,
    pub organization_id: Uuid,
    pub permissions: Vec<String>,
    pub required_permissions: Vec<String>,
    pub frontend: PluginFrontendDto,
}

impl PluginCreateRequest {
    pub fn into_draft(self) -> Result<PluginDraft, ServiceError> {
        Ok(PluginDraft {
            slug: PluginSlug::new(self.slug)?,
            name: PluginName::new(self.name)?,
            description: self.description,
            organization_id: Id::new(self.organization_id),
            permissions: parse_permissions(&self.permissions)?,
            required_permissions: parse_permissions(&self.required_permissions)?,
            frontend: self.frontend.into_domain()?,
        })
    }
}

/// Request body for updating an installed plugin. Every field is optional; an
/// omitted field leaves it untouched. `description` is a double-option so
/// "omitted" and "set to null" (clear it) stay distinguishable, matching
/// `PluginChange`.
#[derive(Debug, Deserialize, utoipa::ToSchema)]
#[schema(example = json!({ "enabled": true }))]
pub struct PluginUpdateRequest {
    pub name: Option<String>,
    #[serde(default, deserialize_with = "deserialize_some")]
    #[schema(value_type = Option<String>)]
    pub description: Option<Option<String>>,
    pub frontend: Option<PluginFrontendDto>,
    pub permissions: Option<Vec<String>>,
    pub required_permissions: Option<Vec<String>>,
    pub enabled: Option<bool>,
}

impl PluginUpdateRequest {
    pub fn into_change(self) -> Result<PluginChange, ServiceError> {
        Ok(PluginChange {
            name: self.name.map(PluginName::new).transpose()?,
            description: self.description,
            frontend: self
                .frontend
                .map(PluginFrontendDto::into_domain)
                .transpose()?,
            permissions: self
                .permissions
                .map(|p| parse_permissions(&p))
                .transpose()?,
            required_permissions: self
                .required_permissions
                .map(|p| parse_permissions(&p))
                .transpose()?,
            enabled: self.enabled,
        })
    }
}

fn deserialize_some<'de, D, T>(deserializer: D) -> Result<Option<T>, D::Error>
where
    T: Deserialize<'de>,
    D: Deserializer<'de>,
{
    Deserialize::deserialize(deserializer).map(Some)
}
