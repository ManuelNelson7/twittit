import React, { useState } from 'react'
import { Container, Row, Col, Button } from "react-bootstrap"
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faSearch, faUsers, faComment } from '@fortawesome/free-solid-svg-icons'

import './SignInSignUp.scss'
import BasicModal from '../../components/modal/BasicModal'
import SignUpForm from '../../components/SignUpForm'
import SignInForm from '../../components/SignInForm'


export default function SignInSingUp(props) {
    const { setRefreshCheckLogin } = props;
    const [showModal, setShowModal] = useState(false);
    const [contentModal, setContentModal] = useState(null);

    const openModal = content => {
        setShowModal(true);
        setContentModal(content);
    };


    return (
        <>
            <Container className="signin-signup" fluid>
                <Row>
                    <LeftComponent />
                    <RightComponent
                        openModal={openModal}
                        setShowModal={setShowModal}
                        setRefreshCheckLogin={setRefreshCheckLogin}
                    />
                </Row>
            </Container>
            <BasicModal show={showModal} setShow={setShowModal} >
                {contentModal}
            </BasicModal>
        </>
    )
}

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
    const { openModal, setShowModal, setRefreshCheckLogin } = props;

    return (
        <Col className='signin-signup__right' xs={6}>
            <div>
                <img src="/img/logo-white.png" alt="" />
                <h2>Mira lo que está pasando en el mundo en este momento</h2>
                <h3>Únete a Twittit hoy mismo</h3>
                <Button
                    variant='primary'
                    onClick={() => openModal(<SignUpForm setShowModal={setShowModal} />)}
                >
                    Regístrate
                </Button>

                <Button
                    variant='outline-primary'
                    onClick={() => openModal(<SignInForm setRefreshCheckLogin={setRefreshCheckLogin} />)}
                >
                    Iniciar sesión
                </Button>
            </div>
        </Col>
    );
}
