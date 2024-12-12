import React from 'react'
import { BrowserRouter as Router , Route,Switch } from 'react-router-dom';
import Signupform from './Components/Signupform';
import Loginform from './Components/Loginform';


const App = () => {
  return (
    <Router>
      <Switch>
        <Route path='/signup' Component={Signupform} />
        <Route path='/login' Component={Loginform} />
      </Switch>
    </Router>
  )
}

export default App
