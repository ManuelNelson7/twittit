import React from 'react'
import { Button } from "react-bootstrap"
import { Link } from "react-router-dom"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import {
    faHome,
    faUser,
    faUsers,
    faPowerOff
} from "@fortawesome/free-solid-svg-icons"
import { logoutApi } from '../../api/auth'

import "./LeftMenu.scss"

const LeftMenu = (props) => {

    const {setRefreshCheckLogin} = props;

    const logout = () => {
        logoutApi();
        setRefreshCheckLogin(true);
    }
    

    return (
        <div className="left-menu">
            <Link to="/"><img className="logo" src="/img/logo-white.png" alt="Twittit" /></Link>

            <Link to="/fassad"><FontAwesomeIcon icon={faUser} />Perfil</Link>
            <Link to="/users"><FontAwesomeIcon icon={faUsers} /> Usuarios</Link>
            <Link to="/logout" onClick={logout}><FontAwesomeIcon icon={faPowerOff} /> Cerrar sesión</Link>

            <Button>
                Nuevo Twitt
            </Button>
        </div>
    )
}

export default LeftMenu
