import { createRootRoute } from "@tanstack/react-router";
import App from "../App";

export const Route = createRootRoute({
  component: Root,
});

function Root() {
  return (
    <div>
      <App />
    </div>
  );
}
