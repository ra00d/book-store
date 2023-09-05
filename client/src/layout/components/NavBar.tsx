import { AppBar, Box, Button, Toolbar } from "@mui/material";
import AbdIcon from "@mui/icons-material/Adb";
import { navItems } from "../../config/routes/navItems";

export default function NavBar() {
  return (
    <AppBar position="static">
      <Toolbar>
        <Button
          startIcon={<AbdIcon />}
          href="/"
          sx={{ color: "white", fontWeight: "bold" }}
        >
          Logo
        </Button>
        <Box>
          {navItems.map((item) => (
            <Button key={item.title} sx={{ color: "white" }} href={item.path}>
              {item.title}
            </Button>
          ))}
        </Box>
      </Toolbar>
    </AppBar>
  );
}
