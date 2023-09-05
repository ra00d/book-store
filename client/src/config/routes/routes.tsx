import { lazy } from "react";
import { TRoute } from "../../types/routes";
import Error404 from "../../pages/404";
import { Navigate } from "react-router-dom";

const LoginPage = lazy(() => import("../../pages/login"));
export const routes: TRoute[] = [
  {
    path: "/login",
    title: "login",
    icon: <></>,
    element: <LoginPage />,
  },
  {
    path: "/sign-up",
    title: "signup",
    icon: <></>,
    element: <LoginPage />,
  },
  {
    path: "/books",
    title: "books",
    icon: <></>,
    element: <LoginPage />,
  },
  {
    path: "/home",
    title: "login",
    icon: <></>,
    element: <LoginPage />,
  },
  {
    path: "/Trend",
    title: "login",
    icon: <></>,
    element: <LoginPage />,
  },
  {
    path: "/404",
    title: "login",
    icon: <></>,
    element: <Error404 />,
  },
  {
    path: "/*",
    title: "login",
    icon: <></>,
    element: <Navigate to="404" />,
  },
];
