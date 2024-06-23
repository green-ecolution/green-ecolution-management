import { TreeDataContextProvider } from "./context/TreeDataContext";
import SideHeader from "./components/SideHeader";
import { Outlet } from "@tanstack/react-router";
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from "./components/ui/resizable";
import { useState } from "react";
import { TooltipProvider } from "./components/ui/tooltip";

function App() {
  const [isCollapsed, setIsCollapsed] = useState(false);
  return (
    <TreeDataContextProvider>
      <TooltipProvider>
        <ResizablePanelGroup direction="horizontal">
          <ResizablePanel
            defaultSize={0}
            minSize={12}
            maxSize={20}
            collapsible
            onCollapse={() => setIsCollapsed(true)}
            onExpand={() => setIsCollapsed(false)}
          >
            <SideHeader open={!isCollapsed} />
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
