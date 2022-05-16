import { Box, Tooltip } from "@mui/material";


type boxInfo = {
    name: string,
    info: string
}

export const NameBox = (props: boxInfo) => {
    return (
        <Tooltip title={props.info}>
            <Box sx={{
                p: 1,
                borderRadius: '25px ',
                backgroundColor: 'primary.dark',
                '&:hover': {
                    backgroundColor: 'primary.main',
                    opacity: [0.9, 0.8, 0.7],
                    cursor: 'pointer'
                },
            }}>
                {props.name}
            </Box>
        </Tooltip>
    );
}