---
slug: settings-plugins
title: Plugins
part: administration
summary: Was ein Plugin ist, wie ein Admin eines installiert, was die beiden Rechtemengen bedeuten, wie der einmalig sichtbare Schlüssel rotiert wird und wie sich die eigene Ansicht eines Plugins öffnen lässt.
routes: ['/settings/plugin', '/settings/plugin/$slug', '/plugin/$slug']
---

> [!NOTE]
> Diese Seite siehst du nur mit der Berechtigung `plugin:read`, und nur, wenn deine
> Instanz die Plugin-Funktion überhaupt eingeschaltet hat. Fehlt dir die Berechtigung
> oder ist die Funktion ausgeschaltet, taucht **Plugins** in deiner Einstellungsnavigation
> gar nicht erst auf. Zum Installieren, Ändern und Deinstallieren brauchst du zusätzlich
> `plugin:create`, `plugin:update` beziehungsweise `plugin:delete`, jeweils in der
> Organisation, der das Plugin gehört.

## Was ein Plugin ist

Ein Plugin ist ein externes System, das Green Ecolution um Daten oder eine eigene
Ansicht erweitert, ohne selbst Teil der Anwendung zu sein. Praktisch bedeutet das: Ein
Plugin bekommt beim Installieren einen eigenen API-Schlüssel und kann darüber Bäume
anlegen, aktualisieren oder löschen, etwa um ein kommunales Baumkataster laufend
einzuspielen. Zusätzlich kann ein Plugin eine eigene Oberfläche mitbringen, die in
Green Ecolution eingebettet erscheint, muss das aber nicht. Ein Plugin registriert sich
dabei nicht selbst; es wird von einer Person mit den passenden Rechten in den
Einstellungen angelegt, ganz wie ein Fahrzeug oder eine Rolle.

Die Übersichtsseite **Plugins** listet jedes installierte Plugin mit Status, Slug,
Organisation und dem Zeitpunkt seines letzten Kontakts. Dieser Zeitpunkt ist keine
manuelle Angabe, sondern aktualisiert sich mit jeder Anfrage, die das Plugin an seine
Schnittstelle stellt, und zeigt dir damit auf einen Blick, ob ein Plugin überhaupt noch
aktiv angebunden ist.

## Ein Plugin installieren

Über **Plugin installieren** öffnest du den Installationsdialog. Slug, Name und
optional eine Beschreibung sowie die Organisation, der das Plugin gehören soll, sind
Pflichtangaben. Der Slug besteht aus Kleinbuchstaben, Ziffern und Bindestrichen und
lässt sich nach dem Anlegen nicht mehr ändern: Er dient zugleich als Herkunftskennung
aller Daten, die das Plugin importiert, und ein nachträglicher Wechsel würde diese
Zuordnung verwaisen lassen.

Darunter legst du fest, ob und wie das Plugin eine eigene Ansicht mitbringt: **Keine
Ansicht** für ein reines Datenplugin, **Extern gehostet** für eine Adresse außerhalb
des Clusters, oder **Im Cluster (proxied)** für einen intern laufenden Dienst. Eine
extern gehostete Ansicht muss eine absolute https-Adresse sein, mit Ausnahme von
localhost für die lokale Entwicklung, und darf nicht auf die Adresse von Green
Ecolution selbst zeigen.

Zum Schluss legst du die beiden Rechtemengen fest, siehe
[Die beiden Rechtemengen](./settings-plugins.md#die-beiden-rechtemengen) weiter unten.
Nach **Installieren** zeigt ein eigener Dialog den frisch erzeugten API-Schlüssel im
Klartext. Das ist der einzige Moment, in dem du ihn zu sehen bekommst: Green Ecolution
speichert nur einen Hashwert davon und kann ihn dir später nicht erneut anzeigen.
Kopiere ihn also sofort dorthin, wo das Plugin ihn zur Anmeldung braucht.

![Der Installationsdialog eines Plugins mit beiden Rechtematrizen](../images/plugin-install.png)

Ein installiertes Plugin lässt sich über seine Detailseite jederzeit umbenennen,
beschreiben, in seiner Ansicht und seinen Rechten anpassen, vorübergehend deaktivieren
oder ganz deinstallieren. Deaktivieren sperrt lediglich die Anmeldung des Plugins an
seiner Schnittstelle, ohne bereits importierte Daten anzutasten; erst **Deinstallieren**
entfernt das Plugin selbst, wobei von ihm bereits angelegte Bäume unangetastet erhalten
bleiben.

## Die beiden Rechtemengen

Der Installationsdialog und die Detailseite zeigen zwei getrennte Rechtematrizen, die
auf den ersten Blick ähnlich aussehen, aber unterschiedliche Fragen beantworten.

**Rechte des Plugins** legt fest, was das Plugin selbst tun darf, wenn es Daten in
Green Ecolution schreibt, etwa Bäume anlegen oder bearbeiten. Diese Rechte darfst du
dem Plugin nur bis zu deinen eigenen Rechten geben; verlangt der Dialog mehr, als du
selbst besitzt, weist Green Ecolution die Eingabe zurück.

**Zugriffsrechte für die Ansicht** legt fest, wer in deiner Organisation die eigene
Oberfläche des Plugins überhaupt öffnen darf, falls es eine hat. Diese Menge übernimmt
beim Anlegen zunächst die Rechte des Plugins, ist davon aber unabhängig: Änderst du sie
eigenständig, bleibt sie ab diesem Zeitpunkt getrennt von den Rechten des Plugins
bestehen. Der Grund für die Trennung ist, dass eine Ansicht mit eigenen
Aktionsknöpfen die Rechte des Plugins an die Person weiterreicht, die sie bedient; wer
also die Zugriffsrechte lockert, sollte bewusst tun, wessen Handlungen er damit
zulässt.

## Den API-Schlüssel rotieren

Vermutest du, dass der Schlüssel eines Plugins bekannt geworden ist, oder wechselt
das Plugin den Betreiber, erneuerst du ihn über **Schlüssel erneuern** auf der
Detailseite. Nach Bestätigung wird der bisherige Schlüssel sofort ungültig, und ein
neuer erscheint, wie beim Installieren, genau einmal im Klartext im selben Dialog.
Bis das Plugin mit dem neuen Schlüssel hinterlegt ist, kann es sich in dieser Zeit
nicht anmelden und keine Daten mehr schreiben.

## Die Ansicht eines Plugins öffnen

Bringt ein Plugin eine externe Ansicht mit, öffnest du sie über **Ansicht öffnen** auf
seiner Detailseite. Green Ecolution bettet sie in einem eigenen, abgeschotteten Rahmen
ein und übergibt ihr deinen Anzeigenamen, die Oberflächensprache und den Slug des
Plugins, aber keinen Zugriff auf deine Anmeldedaten. Ein im Cluster laufendes,
proxiedes Plugin zeigt an dieser Stelle stattdessen einen Hinweis, dass seine Ansicht
hier noch nicht dargestellt werden kann, und ein Plugin ohne eigene Ansicht entsprechend,
dass es keine besitzt.
