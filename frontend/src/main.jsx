import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import CombinedView from './CombinedView';
import Expand from './expand/Expand';
import './index.css';

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <Router>
      <Routes>
        <Route path="/" element={<CombinedView />} />
        <Route path="/expand" element={<Expand />} />
      </Routes>
    </Router>
  </StrictMode>,
);