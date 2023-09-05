import { Box } from "@mui/material";
import { useEffect } from "react";

export default function Error404() {
  useEffect(() => {
    console.log("in login");
  }, []);
  return (
    <Box
      sx={{
        color: "red",
        fontSize: "30px",
        display: "grid",
        alignContent: "center",
      }}
    >
      404
    </Box>
  );
}
