import { useState } from "react";
import SignInSignUp from "./pages/SignInSignUp";

function App() {

  const [user, setUser] = useState({
    user : 'Manuel'
  })

  return (
    <>
      {user ? (
        <SignInSignUp />
      ) : <h1>Not Logeado</h1>}
    </>
  )
}

export default App;
