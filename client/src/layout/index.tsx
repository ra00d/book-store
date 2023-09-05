import { Suspense, useEffect } from "react";
import NavBar from "./components/NavBar";
import { routesStore } from "../config/stores/routes";
import { useRoutes } from "react-router-dom";

export default function Layout() {
  const { getRoutes } = routesStore();
  const route = useRoutes(getRoutes());
  useEffect(() => {
    console.log("routes");
  }, []);
  return (
    <>
      <NavBar />
      <Suspense fallback={<>loading..</>}>{route}</Suspense>
    </>
  );
}
