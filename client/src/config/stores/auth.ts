import { create } from "zustand";
import { devtools, persist } from "zustand/middleware";

export const userStore = create<any>()(
  devtools(
    persist(
      () => ({
        role: "guest",
        permissions: [],
        name: "guest",
        email: "",
        username: "",
      }),
      { name: "user" },
    ),
  ),
);
