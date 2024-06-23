import { MapContainer, Marker, Popup, TileLayer } from "react-leaflet";
import { useTrees } from "../context/TreeDataContext";
import L from "leaflet";

export interface MapProps {
  width?: string;
  height?: string;
}

const markerHtmlStyles = (color: string) => `
  background-color: ${color};
  width: 2rem;
  height: 2rem;
  display: block;
  left: -1rem;
  top: -1rem;
  position: relative;
  border-radius: 3rem 3rem 0;
  transform: rotate(45deg);
  border: 1px solid #FFFFFF`;

export const TreeIcon = (color: string) =>
  L.divIcon({
    className: "my-custom-pin",
    iconAnchor: [0, 24],
    popupAnchor: [0, -36],
    html: `<span style="${markerHtmlStyles(color)}" />`,
  });

const Map = ({ width = "100%", height = "100vh" }: MapProps) => {
  const trees = useTrees();

  const treeMarkers = trees.map((tree) => (
    <Marker
      key={tree.id}
      position={[tree.lat, tree.lng]}
      icon={TreeIcon(
        tree.status === "healthy"
          ? "green"
          : tree.status === "neutral"
            ? "yellow"
            : "red",
      )}
    >
      <Popup>
        <h3 className="font-bold">{tree.name}</h3>
        <span>Status: {tree.status}</span>
        <br />
        <span>Needed Water: {tree.neededWater}l</span>
      </Popup>
    </Marker>
  ));

  return (
    <MapContainer
      className="z-0"
      zoomControl={false}
      style={{ width, height }}
      center={[54.792277136221905, 9.43580607453268]}
      zoom={13}
    >
      <TileLayer
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
      />
      {treeMarkers}
    </MapContainer>
  );
};

export default Map;
