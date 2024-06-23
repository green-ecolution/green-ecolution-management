import { createFileRoute, Link } from "@tanstack/react-router";
import { Button } from "../../components/ui/button";
import { Check, Droplet, Maximize2, PlusCircleIcon, Trees } from "lucide-react";
import { Separator } from "../../components/ui/separator";
import { cn } from "../../lib/utils";
import { useEffect, useMemo, useRef, useState } from "react";
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from "../../components/ui/resizable";
import { Tree, useTrees } from "../../context/TreeDataContext";
import {
  MapContainer,
  Marker,
  Polyline,
  Popup,
  TileLayer,
} from "react-leaflet";
import { TreeIcon } from "../../components/Map";
import { LatLngExpression, Map } from "leaflet";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "../../components/ui/card";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "../../components/ui/tooltip";

export const Route = createFileRoute("/waypoints/")({
  component: Waypoints,
});

interface Waypoint {
  id: string;
  title: string;
  description: string;
  trees?: Tree[];
}

function Waypoints() {
  const trees = useTrees();
  const usedTrees = useMemo(
    () => trees.filter((tree) => tree.status !== "healthy"),
    [trees],
  );
  const waypoints: Waypoint[] = useMemo(
    () => [
      {
        id: "1",
        title: "Einsatzplanung 1",
        description: "Mürwik -> Sonwick -> ...",
        trees: usedTrees.slice(0, 3),
      },
      {
        id: "2",
        title: "Einsatzplanung 2",
        description: "Christiansen Park -> Wassersleben -> ...",
        trees: usedTrees.slice(3, 6),
      },
      {
        id: "3",
        title: "Einsatzplanung 3",
        description: "Zob -> Exe -> ...",
        trees: usedTrees.slice(6, trees.length - 1),
      },
    ],
    [],
  );

  const [selectedWaypoint, setSelectedWaypoint] = useState<Waypoint | null>(
    waypoints[0],
  );
  return (
    <div>
      <ResizablePanelGroup
        direction="horizontal"
        className="min-h-screen items-stretch px-2"
      >
        <ResizablePanel minSize={40}>
          <div className="h-[48px] flex items-center justify-between mx-4">
            <h1 className="font-bold text-xl">Einsatzplanung</h1>

            <div className="flex items-center gap-2">
              <Link to="/waypoints/new">
                <Button variant="default">
                  <PlusCircleIcon className="w-4 h-4" />
                  <span className="ml-2">Einsatzplanung hinzufügen</span>
                </Button>
              </Link>
            </div>
          </div>
          <Separator />
          <div>
            <div className="mx-4 mt-4 flex flex-col gap-2">
              {waypoints.map((waypoint) => (
                <WaypointCard
                  key={waypoint.id}
                  waypoint={waypoint}
                  selected={selectedWaypoint?.id === waypoint.id}
                  onSelect={setSelectedWaypoint}
                />
              ))}
            </div>
          </div>
        </ResizablePanel>
        <ResizableHandle />
        <ResizablePanel>
          {selectedWaypoint ? (
            <WaypointDetails waypoint={selectedWaypoint} />
          ) : (
            <div>Wähle eine Einsatzplanung aus</div>
          )}
        </ResizablePanel>
      </ResizablePanelGroup>
    </div>
  );
}

interface WaypointCardProps {
  waypoint: Waypoint;
  selected?: boolean;
  onSelect?: (waypoint: Waypoint) => void;
}

const WaypointCard = ({ waypoint, selected, onSelect }: WaypointCardProps) => {
  return (
    <Link
      className={cn(
        "flex flex-col items-start gap-2 rounded-lg border p-3 text-left text-sm transition-all hover:bg-accent",
        selected && "bg-muted",
      )}
      to="/waypoints"
      onClick={() => onSelect?.(waypoint)}
    >
      <div className="flex w-full flex-col gap-1">
        <div className="flex items-center">
          <div className="flex items-center gap-2">
            <div className="font-semibold">{waypoint.title}</div>
          </div>
        </div>
        <div className="line-clamp-2 text-xs text-muted-foreground">
          {waypoint.description.substring(0, 300)}
        </div>
      </div>
    </Link>
  );
};

export const WaypointDetails = ({ waypoint }: { waypoint: Waypoint }) => {
  const mapRef = useRef<Map | null>(null);
  const trees = waypoint.trees || [];

  useEffect(() => {
    mapRef.current?.flyTo([trees[0].lat, trees[0].lng]);
  }, [waypoint]);

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

  const lineToTree: LatLngExpression[] = useMemo(
    () => trees.map((tree) => [tree.lat, tree.lng]),
    [trees],
  );

  const onTreeClick = (tree: Tree) => {
    mapRef.current?.flyTo([tree.lat, tree.lng]);
  };

  return (
    <div>
      <div className="h-[48px] flex items-center justify-between mx-4">
        <Tooltip>
          <TooltipTrigger>
            <Button variant="ghost" size="icon">
              <Maximize2 className="w-5 h-5" />
            </Button>
          </TooltipTrigger>
          <TooltipContent>Vollbild</TooltipContent>
        </Tooltip>

        <div className="flex items-center gap-2">
          <div className="h-8 w-8 bg-grey-100 rounded-xl"></div>
          <div className="h-8 w-8 bg-grey-100 rounded-xl"></div>
          <div className="h-8 w-8 bg-grey-100 rounded-xl"></div>
          <div className="h-8 w-8 bg-grey-100 rounded-xl"></div>
        </div>
      </div>
      <Separator />
      <div className="m-2 grid grid-cols-2 gap-4">
        <Card>
          <CardHeader>
            <CardTitle>Alle Bäume</CardTitle>
            <CardDescription>
              Klicke auf einen Eintrag um den Baum in der Karte zu finden
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="flex gap-2 flex-col">
              {trees.map((tree, index) => (
                <div
                  className="flex border rounded items-center p-2 gap-2 cursor-pointer hover:bg-muted transition-all justify-between"
                  key={tree.id}
                  onClick={() => onTreeClick(tree)}
                >
                  <div className="flex items-center gap-2">
                    <Trees className="w-8 h-8" />
                    <div>
                      <h4 className="font-bold">{tree.name}</h4>
                      <p>Status: {tree.status}</p>
                    </div>
                  </div>
                  <div className="">
                    {index === 0 ? (
                      <Tooltip>
                        <TooltipTrigger>
                          <Check className="w-6 h-6" />
                        </TooltipTrigger>
                        <TooltipContent>
                          <h4 className="font-bold">Fortschritt</h4>
                          <p>
                            Dieser Baum wurde als erledigt markiert und benötigt
                            keine weitere Pflege
                          </p>
                        </TooltipContent>
                      </Tooltip>
                    ) : (
                      <Tooltip>
                        <TooltipTrigger>
                          <Droplet className="w-6 h-6" />
                        </TooltipTrigger>
                        <TooltipContent>
                          <h4 className="font-bold">Fortschritt</h4>
                          <p>
                            Dieser Baum wurde noch nicht als erledigt markiert
                          </p>
                        </TooltipContent>
                      </Tooltip>
                    )}
                  </div>
                </div>
              ))}
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Daten</CardTitle>
            <CardDescription>Informationen zur Einsatzplanung</CardDescription>
          </CardHeader>
          <CardContent>
            <div className="flex flex-col gap-2">
              <div className="flex items-center justify-between">
                <h4 className="font-bold">Name</h4>
                <p>{waypoint.title}</p>
              </div>
              <div className="flex items-center justify-between">
                <h4 className="font-bold">Benötigte Wassermenge</h4>
                <code className="relative rounded bg-muted px-[0.3rem] py-[0.2rem] font-mono text-sm font-semibold">
                  {trees.reduce(
                    (acc, tree) => acc + (tree.neededWater || 0),
                    0,
                  )}{" "}
                  Liter
                </code>
              </div>
              <div className="flex items-center justify-between">
                <h4 className="font-bold">Geschäzte Zeit</h4>
                <code className="relative rounded bg-muted px-[0.3rem] py-[0.2rem] font-mono text-sm font-semibold">
                  2 Stunden
                </code>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>

      <Card className="m-2">
        <CardContent className="p-0 m-0">
          <MapContainer
            style={{ width: "100%", height: "400px" }}
            ref={mapRef}
            center={[54.792277136221905, 9.43580607453268]}
            zoom={13}
            scrollWheelZoom
            zoomControl={false}
          >
            <TileLayer
              attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
              url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            />
            {treeMarkers}
            <Polyline positions={lineToTree} />
          </MapContainer>
        </CardContent>
      </Card>
    </div>
  );
};
