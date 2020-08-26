import * as serviceWorker from "./serviceWorker";

import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { store } from "./store";

import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";

import './index.css'

import Home from "./pages/home";
import App from "./pages/app";

ReactDOM.render(
  <Provider store={store}>
    <div className="index">
      <Router>
        <Switch>

          <Route path="/" exact>
            <Home />
          </Route>

          <Route path="/app" exact>
            <App />
          </Route>

        </Switch>
      </Router>
    </div>
  </Provider>,
  document.getElementById("root")
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
