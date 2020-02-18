import * as React from 'react'
import {Link} from "react-router-dom";
import {Button} from "../../App/App.components/Button/Button.controller";

import {HomeButton, NotFoundStyled, NotFoundText} from "./NotFound.style";

export const NotFound = () => (
    <NotFoundStyled>
        <img alt="Unplugged" height="50" src="/icons/unplugged.svg"/> Oops
        <NotFoundText>There isn't anything here...</NotFoundText>
        <HomeButton>
            <Link to="/">
                <Button text="Go to Home" icon="cards" />
            </Link>
        </HomeButton>
    </NotFoundStyled>
)
