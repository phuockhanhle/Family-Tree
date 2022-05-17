import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { DashBoardPage } from './pages/dashboard/DashBoardPage';
import { LoginPage } from './pages/login/LoginPage';

export const App = () => {
  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
            <Route path="/" element={<LoginPage/>} />
            <Route path="/dashboard" element={<DashBoardPage/>} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}