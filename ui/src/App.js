import "bootstrap/dist/css/bootstrap.min.css";
import { Route, Switch } from "react-router";
import { BrowserRouter } from "react-router-dom";
import "./App.css";
import Home from "./common/Home";
import Seat from "./common/Seat";
import PokerTable from "./common/Table";

function App() {
  return (
    <BrowserRouter>
      <Switch>
        <Route path="/" exact={true}>
          <Home />
        </Route>

        <Route path="/table" exact={true}>
          <PokerTable />
        </Route>

        <Route path="/player/:id" exact={true}>
          <Seat />
        </Route>
      </Switch>
    </BrowserRouter>
  );
}

export default App;
