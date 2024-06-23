import { Menu } from "lucide-react";
import { Button } from "./ui/button";

export interface HeaderProps {}

const MapHeader = ({}: HeaderProps) => {
  return (
    <div className="z-50 absolute top-4 left-4 w-[350px] h-12 bg-white rounded shadow-lg">
      <div className="flex justify-between items-center h-full mx-2">
        <div className="flex items-center">
          <Button variant="ghost" size="icon">
            <Menu className="w-6 h-6" />
          </Button>
        </div>
        <div className="flex items-center gap-2">
          <div className="flex items-center">
            <div className="h-8 w-8 bg-grey-100 rounded-xl"></div>
          </div>
          <div className="flex items-center">
            <div className="h-8 w-8 bg-grey-100 rounded-xl"></div>
          </div>
          <div className="flex items-center">
            <div className="h-8 w-8 bg-grey-100 rounded-xl"></div>
          </div>
          <div className="flex items-center">
            <div className="h-8 w-8 bg-grey-100 rounded-xl"></div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default MapHeader;
