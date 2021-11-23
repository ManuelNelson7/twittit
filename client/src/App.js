import { useState, useEffect } from "react";
import SignInSignUp from "./pages/SignInSignUp";
import { ToastContainer } from "react-toastify"
import { AuthContext } from "./utils/contexts";
import { isUserLoggedApi } from "./api/auth";
import Routing from "./routes/Routing";

function App() {

  const [user, setUser] = useState(null)
  const [loadUser, setLoadUser] = useState(false)
  const [refreshCheckLogin, setRefreshCheckLogin] = useState(false)

  useEffect(() => {
    setUser(isUserLoggedApi());
    setRefreshCheckLogin(false);
    setLoadUser(true);
  }, [refreshCheckLogin]);

  if(!loadUser) return null;

  return (
    <AuthContext.Provider value={user}>
      {user? (
        <h1>Logeado</h1>
      ) : (
        <SignInSignUp setRefreshCheckLogin={setRefreshCheckLogin} />
      )}
      <ToastContainer
        position="top-right"
        autoclose={5000}
        hideprogressbar
        newestontop={false}
        closeoncLick
        rtl={false}
        pauseonvisibilitychange
        draggable
        pauseonhover="true"
      />
    </AuthContext.Provider>
  )
}

export default App;
