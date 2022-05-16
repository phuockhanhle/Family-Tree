import { Box, Tooltip } from "@mui/material";
import { useEffect, useState } from "react";


type boxInfo = {
    name: string,
    info: string,
    person: any
}

export const NameBox = (props: boxInfo) => {
    const [show, setShow] = useState(true);
    const hideID = (name: string) => {
        if (name === props.person.id.name) {
            setShow(false);
        }
    }  
    return (
        <Tooltip title={props.info}>
            <Box sx={{
                visibility: show ? 'visible' : 'hidden',
                p: 1,
                borderRadius: '25px ',
                backgroundColor: 'primary.dark',
                '&:hover': {
                    backgroundColor: 'primary.main',
                    opacity: [0.9, 0.8, 0.7],
                    cursor: 'pointer'
                },
            }} onClick={() => hideID(props.name)}>
                {props.name}
            </Box>
        </Tooltip>
    );
}