import { BrowserRouter, Routes, Route } from "react-router";
import WelcomePage from "./components/WelcomePage";
import Login from "./components/Play";
function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route exact path="/" element={<WelcomePage />} />
        <Route exact path="/play" element={<Login />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
