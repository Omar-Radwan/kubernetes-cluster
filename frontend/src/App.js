import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Mongo from "./pages/Mongo";
import Redis from "./pages/Redis";
import Root from "./pages/Root";

function App() {
  return (
    <div className="App">
      <Router>
        <Routes>

          <Route path="/" element={<Root />} />
          <Route path="/redis" element={<Redis />} />
          <Route path="/mongo" element={<Mongo />} />

        </Routes>
      </Router>
    </div>
  );
}

export default App;
