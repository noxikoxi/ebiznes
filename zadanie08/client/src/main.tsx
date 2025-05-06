import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom';
import LoginPage from "./pages/LoginPage.tsx";
import HelloPage from "./pages/HelloPage.tsx";
import RegisterPage from "./pages/RegisterPage.tsx";
import * as React from "react";

createRoot(document.getElementById('root')!).render(
  <StrictMode>
      <Router>
          <Routes>
              <Route path='/login' element={<LoginPage/>}/>
              <Route path='/hello' element={<HelloPage/>}/>
              <Route path='/register' element={<RegisterPage/>}/>
          </Routes>
      </Router>
  </StrictMode>,
)
