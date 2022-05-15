import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { DashBoardPage } from './components/pages/dashboard/DashBoardPage';
import { LoginPage } from './components/pages/login/LoginPage';

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