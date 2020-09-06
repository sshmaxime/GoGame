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

import HomePage from "./pages/home";

ReactDOM.render(
  <Provider store={store}>
    <Router>
      <Switch>

        <Route path="/" exact>
          <HomePage />
        </Route>

      </Switch>
    </Router>
  </Provider>,
  document.getElementById("root")
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
