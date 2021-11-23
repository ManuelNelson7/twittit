import Home from "../pages/Home/Home";
import Error404 from "../pages/Error404/Error404";

export default [
    {
        path: "/",
        exact: true,
        page: Home
    },
    {
        path:"*",
        page: Error404
    }
]