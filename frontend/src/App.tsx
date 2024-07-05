import { TreeDataContextProvider } from "./context/TreeDataContext";
import SideHeader from "./components/SideHeader";
import { Outlet } from "@tanstack/react-router";
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from "./components/ui/resizable";
import { TooltipProvider } from "./components/ui/tooltip";
import useSidePanelStore from "./store/sidePanelStore";
import { useEffect, useRef } from "react";
import { ImperativePanelHandle } from "react-resizable-panels";

function App() {
  const panelRef = useRef<ImperativePanelHandle>(null);
  const { isOpen, close, open } = useSidePanelStore((state) => ({
    isOpen: state.isOpen,
    close: state.close,
    open: state.open,
  }));

  useEffect(() => {
    if (isOpen) {
      panelRef.current?.expand(20);
    } else {
      panelRef.current?.collapse();
    }
  }, [isOpen]);

  return (
    <TreeDataContextProvider>
      <TooltipProvider>
        <ResizablePanelGroup direction="horizontal">
          <ResizablePanel
            ref={panelRef}
            defaultSize={0}
            minSize={12}
            maxSize={20}
            collapsible
            onCollapse={() => close()}
            onExpand={() => open()}
          >
            <SideHeader open={isOpen} />
          </ResizablePanel>
          <ResizableHandle withHandle />
          <ResizablePanel>
            <Outlet />
          </ResizablePanel>
        </ResizablePanelGroup>
      </TooltipProvider>
    </TreeDataContextProvider>
  );
}

export default App;
