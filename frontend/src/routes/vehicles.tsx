import { createFileRoute } from "@tanstack/react-router";
import { Separator } from "../components/ui/separator";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "../components/ui/table";
import { Button } from "../components/ui/button";
import { Edit, PlusCircleIcon, Trash } from "lucide-react";

export const Route = createFileRoute("/vehicles")({
  component: Vehicles,
});

const vehicles = [
  {
    title: "LF 10",
    licensePlate: "FL-TB-1235",
    type: "Wasserfahrzeug",
    location: "Klärwerk",
    status: "Verfügbar",
  },
  {
    title: "LF 20",
    licensePlate: "FL-TB-1235",
    type: "Wasserfahrzeug",
    location: "TBZ Standort",
    status: "Verfügbar",
  },
  {
    title: "LF 10",
    licensePlate: "FL-TB-1235",
    type: "Wasserfahrzeug",
    location: "TBZ Standort",
    status: "Verfügbar",
  },
  {
    title: "LF 20",
    licensePlate: "FL-TB-1235",
    type: "Pritschenwagen",
    location: "Klärwerk",
    status: "Nicht verfügbar",
  },
];

function Vehicles() {
  return (
    <div>
      <div className="h-[48px] flex items-center justify-between mx-4">
        <h1 className="font-bold text-xl">Fahrzeuge</h1>

        <div className="flex items-center gap-2">
          <div className="h-8 w-8 bg-grey-100 rounded-xl"></div>
          <div className="h-8 w-8 bg-grey-100 rounded-xl"></div>
          <div className="h-8 w-8 bg-grey-100 rounded-xl"></div>
          <Button variant="default">
            <PlusCircleIcon className="w-4 h-4" />
            <span className="ml-2">Fahrzeug hinzufügen</span>
          </Button>
        </div>
      </div>
      <Separator />

      <div className="p-4">
        <div className="flex justify-end items-center"></div>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead className="w-[100px]">Beziechnung</TableHead>
              <TableHead>Kennzeichen</TableHead>
              <TableHead>Typ</TableHead>
              <TableHead>Standort</TableHead>
              <TableHead>Status</TableHead>
              <TableHead className="text-right">Aktion</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {vehicles.map((vehicle) => (
              <TableRow key={vehicle.title}>
                <TableCell className="font-medium">{vehicle.title}</TableCell>
                <TableCell>{vehicle.licensePlate}</TableCell>
                <TableCell>{vehicle.type}</TableCell>
                <TableCell>{vehicle.location}</TableCell>
                <TableCell>{vehicle.status}</TableCell>
                <TableCell className="text-right">
                  <Button variant="ghost" size="icon">
                    <Edit className="w-4 h-4" />
                  </Button>
                  <Button variant="ghost" size="icon">
                    <Trash className="w-4 h-4" />
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
    </div>
  );
}
