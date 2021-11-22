import React, { useState } from 'react'
import { Container, Row, Col, Button } from "react-bootstrap"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faSearch, faUsers, faComment } from '@fortawesome/free-solid-svg-icons'
import BasicModal from '../../components/modal/BasicModal/'


import './SignInSignUp.scss'

const SignInSignUp = () => {
    const [showModal, setShowModal] = useState(false)
    const [contentModal, setContentModal] = useState(null)

    const openModal = content => {
        setShowModal(true);
        setContentModal(content);
    };

    return (
        <>
            <Container className='signin-signup' fluid>
                <Row>
                    <LeftComponent />
                    <RightComponent openModal={openModal} setShowModal={setShowModal} />
                </Row>
            </Container>
            <BasicModal show={showModal} setShow={setShowModal} >
                {contentModal}
            </BasicModal>
        </>
    )
}

export default SignInSignUp

const LeftComponent = () => {
    return (
        <Col className='signin-signup__left' xs={6}>
            <img src="/img/logo.png" alt="Twittit" />
            <div>
                <h2><FontAwesomeIcon icon={faSearch} />Sigue lo que te interesa</h2>
                <h2><FontAwesomeIcon icon={faUsers} /> Entérate de qué está sucediendo</h2>
                <h2><FontAwesomeIcon icon={faComment} /> Únete a la conversación</h2>
            </div>
        </Col>
    )
}

const RightComponent = (props) => {
    const { openModal, setShowModal } = props;

    return (
        <Col className='signin-signup__right' xs={6}>
            <div>
                <img src="/img/logo-white.png" alt="" />
                <h2>Mira lo que está pasando en el mundo en este momento</h2>
                <h3>Únete a Twittit hoy mismo</h3>
                <Button
                    variant='primary'
                    onClick={() => openModal(<h2>Formulario de Registro</h2>)}
                >
                    Regístrate
                </Button>

                <Button
                    variant='outline-primary'
                    onClick={() => openModal(<h2>Formulario de Login</h2>)}
                >
                    Iniciar sesión
                </Button>
            </div>
        </Col>
    )
}
