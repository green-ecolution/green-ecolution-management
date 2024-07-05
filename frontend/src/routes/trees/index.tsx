import { createFileRoute } from "@tanstack/react-router";
import { Separator } from "../../components/ui/separator";

export const Route = createFileRoute("/trees/")({
  component: Trees,
});

function Trees() {
  return (
    <div className="h-screen">
      <div className="h-[48px] flex items-center justify-between mx-4">
        <h1 className="font-bold text-xl">Übersicht der Bäume</h1>

        <div className="flex items-center gap-2">
        </div>
      </div>

      <Separator />

      <div className="p-4">
        <div>
        </div>
      </div>
    </div>
  );
}


