import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { NameBox } from "../../boxes/nameBox/NameBox";
import Typography from "@mui/material/Typography";
import { Grid, Box, Tooltip, Button } from "@mui/material";


type dashBoardProps = {
    person: any
}

export const DashBoard = (props: dashBoardProps) => {
    const familyTree = [];
    // add parents
    for (let i = 0; i < props.person.parents.total; i++) {
        familyTree.push(
            <Grid item>
                <NameBox name={props.person.parents.details[i].name} info={props.person.parents.details[i].gender} />
            </Grid>
        )
    }
    // add self
    familyTree.push(
        <Grid item>
            <NameBox name={props.person.id.name} info={props.person.id.gender} />
        </Grid >
    )
    // add siblings
    for (let i = 0; i < props.person.siblings.total; i++) {
        familyTree.push(
            <Grid item>
                <NameBox name={props.person.siblings.details[i].name} info={props.person.siblings.details[i].gender} />
            </Grid >
        )
    }
    // add children
    for (let i = 0; i < props.person.children.total; i++) {
        familyTree.push(
            <Grid item>
                <NameBox name={props.person.children.details[i].name} info={props.person.children.details[i].gender} />
            </Grid >
        )
    }

    return (
        <Grid container spacing={1}>
            {familyTree}
        </Grid>
    );
}
