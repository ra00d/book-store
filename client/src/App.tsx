import { BrowserRouter } from "react-router-dom";
import Layout from "./layout";
function App() {
  return (
    <BrowserRouter basename="/">
      <Layout />
    </BrowserRouter>
  );
}

export default App;
