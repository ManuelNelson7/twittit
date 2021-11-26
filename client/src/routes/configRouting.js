import Home from "../pages/Home/Home";
import Error404 from "../pages/Error404/Error404";
import User from "../pages/User/";

export default [
    {
        path: "/:id",
        exact: true,
        page: <User />

    },
    {
        path: "/",
        exact: true,
        page: <Home />
    },
    {
        path:"*",
        page: <Error404 />
    }
]