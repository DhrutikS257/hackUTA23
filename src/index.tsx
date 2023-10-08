/* @refresh reload */
import { render } from 'solid-js/web'
import { Router,Routes, Route } from "@solidjs/router";

import './styles/index.css';
import Login from './Login';
import Signup from './Signup';
import Dashboard from './Dashboard';

const root = document.getElementById('root')

render(() => (
    <Router>
      <Routes>
        <Route path="/" component={Login} />
        <Route path="/signup" component={Signup} />
        <Route path="/dashboard" component={Dashboard} />
      </Routes>
    </Router>
), root!)
