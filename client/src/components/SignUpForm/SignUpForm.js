import React, { useState } from 'react'
import './SignUpForm.scss'
import { Row, Col, Form, Button, Spinner } from "react-bootstrap"
import { values, size } from "lodash"
import { toast } from "react-toastify"
import { isEmailValid } from '../../utils/validations'
import { signUpApi } from '../../api/auth'

const SignUpForm = (props) => {
    const { setShowModal } = props;
    const [formData, setFormData] = useState(initialFormValue());
    const [signUpLoading, setSignUpLoading] = useState(false);

    const onSubmit = e => {
        e.preventDefault();
        console.log(FormData)

        let validCount = 0;
        values(formData).some(value => {
            value && validCount++
            return null
        });

        console.log(validCount);

        if (size(formData) !== validCount) {
            toast.warning("Completa todos los campos del formulario")
        } else {
            if (!isEmailValid(formData.email)) {
                toast.warning("El Email es inválido")
            } else if (formData.password !== formData.repeatPassword) {
                toast.warning("Las contraseñas deben ser iguales")
            } else if (size(formData.password) < 6) {
                toast.warning("La contraseña debe tener al menos 6 caracteres")
            } else {
                setSignUpLoading(true);
                signUpApi(formData).then(response => {
                    if (response.code) {
                        toast.warning(response.message)
                    } else {
                        toast.success("Te registraste exitosamente!");
                        setShowModal(false);
                        setFormData(initialFormValue());
                    }
                })
                    .catch(() => {
                        toast.error("Error del servidor, inténtelo más tarde")
                    })
                    .finally(() => {
                        setSignUpLoading(false);
                    })
            }
        }

    };

    const onChange = e => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    return (
        <div className="sign-up-form">
            <h2>Crea tu cuenta</h2>
            <Form onSubmit={onSubmit} onChange={onChange}>

                <Form.Group className="form-group">
                    <Row>
                        <Col>
                            <Form.Control
                                type="text"
                                placeholder="Nombre"
                                name="name"
                                defaultValue={formData.name}
                            />
                        </Col>
                        <Col>
                            <Form.Control
                                type="text"
                                placeholder="Apellido"
                                name="lastName"
                                defaultValue={formData.lastName}
                            />
                        </Col>
                    </Row>
                </Form.Group>

                <Form.Group className="form-group">
                    <Form.Control
                        type="email"
                        placeholder="Email"
                        name="email"
                        defaultValue={formData.email}
                    />
                </Form.Group>

                <Form.Group className="form-group">
                    <Row>
                        <Col>
                            <Form.Control
                                type="password"
                                placeholder="Contraseña"
                                name="password"
                                defaultValue={formData.password}
                            />
                        </Col>
                        <Col>
                            <Form.Control
                                type="password"
                                placeholder="Confirmar contraseña"
                                name="repeatPassword"
                                defaultValue={formData.repeatPassword}
                            />
                        </Col>
                    </Row>
                </Form.Group>

                <Button variant="primary" type="submit">
                    {!signUpLoading ? "Registrarse" : <Spinner animation="border" />}
                </Button>
            </Form>

        </div>
    )
}

export default SignUpForm


function initialFormValue() {
    return {
        name: "",
        lastName: "",
        email: "",
        password: "",
        repeatPassword: ""
    };
}
