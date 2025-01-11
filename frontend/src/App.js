import { BrowserRouter, Routes, Route } from "react-router";
import WelcomePage from "./components/WelcomePage";
function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route exact path="/" element={<WelcomePage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
