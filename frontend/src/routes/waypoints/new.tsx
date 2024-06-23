import { createFileRoute, Link } from "@tanstack/react-router";
import { MapContainer, Marker, Polyline, TileLayer } from "react-leaflet";
import { Tree, useTrees } from "../../context/TreeDataContext";
import { TreeIcon } from "../../components/Map";
import {
  Map,
  Marker as LMarker,
  LeafletMouseEvent,
  LatLng,
  LatLngExpression,
} from "leaflet";
import MapHeader from "../../components/MapHeader";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "../../components/ui/card";
import { ForwardedRef, forwardRef, useMemo, useRef, useState } from "react";
import { Button } from "../../components/ui/button";
import { Eye, Trash } from "lucide-react";
import { ScrollArea } from "../../components/ui/scroll-area";

export const Route = createFileRoute("/waypoints/new")({
  component: NewWaypoint,
});

function NewWaypoint() {
  const trees = useTrees();
  const [lastSelectedTree, setLastSelectedTree] = useState<Tree>();
  const [selectedTrees, setSelectedTrees] = useState<Tree[]>([]);
  const mapRef = useRef<Map>(null);

  const markerEvenetHandlers = useMemo(
    () => ({
      click: (e: LeafletMouseEvent) => {
        const marker = e.target;
        const tree = trees.find(
          (tree) =>
            tree.lat === marker.getLatLng().lat &&
            tree.lng === marker.getLatLng().lng,
        );
        if (tree) {
          setLastSelectedTree(tree);
          setSelectedTrees((prev) => {
            if (prev.includes(tree)) {
              return prev.filter((t) => t !== tree);
            } else {
              return [...prev, tree];
            }
          });
        }
      },
    }),
    [],
  );

  const treePolyline: LatLngExpression[] = selectedTrees.map((tree) => [
    tree.lat,
    tree.lng,
  ]);

  const onTreeClick = (tree: Tree) => {
    setLastSelectedTree(tree);
    mapRef.current?.flyTo([tree.lat, tree.lng]);
  };

  const onTreeRemove = (tree: Tree) => {
    setSelectedTrees((prev) => prev.filter((t) => t !== tree));
  };

  const treeMarkers = trees.map((tree) => (
    <Marker
      eventHandlers={markerEvenetHandlers}
      key={tree.id}
      position={[tree.lat, tree.lng]}
      icon={TreeIcon(
        tree.status === "healthy"
          ? "green"
          : tree.status === "neutral"
            ? "yellow"
            : "red",
      )}
    ></Marker>
  ));
  return (
    <div className="relative">
      <MapHeader />
      <NewWaypointControl
        selection={selectedTrees}
        onTreeClick={onTreeClick}
        onTreeRemove={onTreeRemove}
        lastSelectedTree={lastSelectedTree}
      />
      <NewWaypointMap ref={mapRef}>
        {treeMarkers}
        <Polyline pathOptions={{ color: "blue" }} positions={treePolyline} />
      </NewWaypointMap>
    </div>
  );
}

interface NewWaypointControlProps {
  selection: Tree[];
  lastSelectedTree?: Tree;
  onTreeClick: (tree: Tree) => void;
  onTreeRemove: (tree: Tree) => void;
}

const NewWaypointControl = ({
  selection,
  onTreeClick,
  onTreeRemove,
  lastSelectedTree
}: NewWaypointControlProps) => {
  return (
    <div className="absolute z-50 top-4 right-4">
      <Card className="">
        <CardHeader>
          <CardTitle>Neue Einsatzplanung</CardTitle>
          <CardDescription>
            Wähle die Bäume aus, die gegossen werden sollen um diese in der
            Planung zu berücksichtigen
          </CardDescription>
        </CardHeader>
        <CardContent>
          {selection.length === 0 ? (
            <p className="text-muted-foreground">
              Wähle Bäume aus, um diese in die Planung aufzunehmen
            </p>
          ) : (
            <div className="w-full">
              {selection.map((tree) => (
                <div
                  key={tree.id}
                  className="flex items-center gap-2 justify-between"
                >
                  <div className="flex items-center gap-2">
                    <div
                      onClick={() => onTreeClick(tree)}
                      className="h-8 w-8 bg-grey-100 rounded-xl cursor-pointer hover:scale-105 transition-transform"
                      style={{
                        backgroundColor:
                          tree.status === "healthy"
                            ? "green"
                            : tree.status === "neutral"
                              ? "yellow"
                              : "red",
                      }}
                    ></div>
                    <span>{tree.name}</span>
                  </div>
                  <div>
                    <Button
                      variant="ghost"
                      size="icon"
                      onClick={() => onTreeClick(tree)}
                    >
                      <Eye className="w-4 h-4" />
                    </Button>
                    <Button
                      variant="ghost"
                      size="icon"
                      onClick={() => onTreeRemove(tree)}
                    >
                      <Trash className="w-4 h-4" />
                    </Button>
                  </div>
                </div>
              ))}
            </div>
          )}
        </CardContent>
        <CardFooter className="flex items-center justify-end gap-2">
          <Link to="/waypoints">
            <Button variant="secondary">Abbrechen</Button>
          </Link>

          <Link to="/waypoints">
            <Button variant="default" disabled={selection.length <= 0}>
              Planung erstellen
            </Button>
          </Link>
        </CardFooter>
      </Card>

      <div className="mt-4">
        <TreeInfo tree={lastSelectedTree} />
      </div>
    </div>
  );
};

interface MapProps extends React.PropsWithChildren { }

const NewWaypointMap = forwardRef(
  ({ children }: MapProps, ref: ForwardedRef<Map>) => {
    return (
      <MapContainer
        ref={ref}
        className="z-0"
        zoomControl={false}
        style={{ width: "100%", height: "100vh" }}
        center={[54.792277136221905, 9.43580607453268]}
        zoom={13}
      >
        <TileLayer
          attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        />
        {children}
      </MapContainer>
    );
  },
);

const TreeInfo = ({ tree }: { tree?: Tree }) => {
  if (!tree) return <></>;
  return (
    <Card>
      <CardHeader>
        <CardTitle>{tree.name}</CardTitle>
        <CardDescription>Informationen zum ausgewählten Baum</CardDescription>
      </CardHeader>
      <CardContent>
        <div className="grid grid-cols-2 gap-2">
          <img className="max-h-full max-w-full border" height="360px" width="250px" src={tree.image} alt={tree.name} />
          <ScrollArea>
            <p>
              {tree.status === "healthy"
                ? "Der Baum ist gesund"
                : tree.status === "neutral"
                  ? "Der Baum benötigt Wasser"
                  : "Der Baum benötigt dringend Wasser"}
            </p>
            <p>Benötigte Wassermenge: {tree.neededWater}l</p>
            <p>Letzte Bewässerung: vor 1 Tag</p>
            <p>Letzte Kontrolle: vor 2 Wochen</p>
            <p>Standort: {tree.lat}, {tree.lng}</p>
            <p>Adresse: Musterstr.</p>
            <p>Art: {tree.treeType}</p>
            <p>Alter: {tree.treeAge}</p>
            <p>Größe: {tree.treeHeight}m</p>
            <p>Umfang: {tree.treeCrown}</p>
            <p>Kronendurchmesser: {tree.treeDiameter}</p>
            <p>Bodenart: schluffig</p>
            <p>Bodenfeuchte: 30%</p>
            <p>Verkehrssicherheit: sicher</p>
            <p>Wurzelraum: {tree.treeRootStructure}</p>
          </ScrollArea>
        </div>
      </CardContent>
    </Card>
  );
};
