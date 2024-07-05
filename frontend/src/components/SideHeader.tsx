import {
  Book,
  Calendar,
  MapPin,
  Settings,
  Trees,
  Truck,
  User,
  Waypoints,
} from "lucide-react";
import { Separator } from "./ui/separator";
import { Link } from "@tanstack/react-router";

export interface SideHeaderProps {
  open: boolean;
}

const SideHeader = ({ open }: SideHeaderProps) => {
  if (!open) {
    return <></>;
  }

  const items = [
    {
      title: "Karte",
      icon: <MapPin className="w-5 h-5" />,
      to: "/",
    },
    {
      title: "Einsatzplanung",
      icon: <Waypoints className="w-5 h-5" />,
      to: "/waypoints",
    },
    {
      title: "Bäume",
      icon: <Trees className="w-5 h-5" />,
      to: "/trees",
    },
    {
      title: "Team",
      icon: <User className="w-5 h-5" />,
      to: "/team",
    },
    {
      title: "Kalender",
      icon: <Calendar className="w-5 h-5" />,
      to: "/calendar",
    },
    {
      title: "Dokumete",
      icon: <Book className="w-5 h-5" />,
      to: "/documents",
    },
    {
      title: "Fahrzeuge",
      icon: <Truck className="w-5 h-5" />,
      to: "/vehicles",
    },
    {
      title: "Einstellungen",
      icon: <Settings className="w-5 h-5" />,
      to: "/settings",
    },
  ];

  return (
    <div className="z-50 h-screen bg-white rounded shadow">
      <div className="flex flex-col h-full">
        <div className="flex justify-between items-center h-12 mx-2">
          <div className="flex items-center">
            <img className="h-6" src="/logo-large-color.svg" alt="logo" />
            <h1 className="font-bold text-xl ml-3">Green Ecolution</h1>
          </div>
        </div>
        <Separator />
        <div className="ml-2 flex flex-col gap-2">
          {items.map((item) => (
            <Link
              key={item.title}
              to={item.to}
              className="flex items-center gap-2 p-2 hover:bg-muted hover:cursor-pointer"
            >
              <div>{item.icon}</div>
              <div>{item.title}</div>
            </Link>
          ))}
        </div>
      </div>
    </div>
  );
};

export default SideHeader;
