import { useCallback, useState } from "react";
import { useNavigate } from "react-router-dom";
import { NameBox } from "../../boxes/nameBox/NameBox";
import Typography from "@mui/material/Typography";
import { Grid, Box, Tooltip, Button } from "@mui/material";


type familyProps = {
    person: any
}

export const FamilyBox = (props: familyProps) => {
    const familyTree: any = [];
    const [person, setPerson] = useState(props.person)

    // add parents
    for (let i = 0; i < person.parents.total; i++) {
        familyTree.push(
            <NameBox name={person.parents.details[i].name} info={person.parents.details[i].gender} person={person}/>
        )
    }
    // add self
    familyTree.push(
        <NameBox name={person.id.name} info={person.id.gender} person={person}/>
    )
    // add siblings
    for (let i = 0; i < person.siblings.total; i++) {
        familyTree.push(
            <NameBox name={person.siblings.details[i].name} info={person.siblings.details[i].gender} person={person}/>
        )
    }
    // add children
    for (let i = 0; i < person.children.total; i++) {
        familyTree.push(
            <NameBox name={person.children.details[i].name} info={person.children.details[i].gender} person={person}/>
        )
    }

    return (
        <Grid container spacing={1}>
            {familyTree.map((member: any) => {
                return (<Grid item>
                    {member}
                </Grid>)
            })}
        </Grid>
    );
}
