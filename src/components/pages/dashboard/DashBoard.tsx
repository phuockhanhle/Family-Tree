import { useState } from "react";
import { Container, Grid, Divider, Typography } from "@material-ui/core";
import { useNavigate } from "react-router-dom";
import { useStyles } from "../../utils/defaultStyles";
import { NameBox } from "../../boxes/nameBox/NameBox";
import { MenuDropDown } from "../../boxes/menuDropDown/MenuDropDown";


type dashBoardProps = {
    parent: any
    child: any
}

export const DashBoard = (props: dashBoardProps) => {
    const classes = useStyles();
    const navigate = useNavigate();
    const [anchorEl, setAnchorEl] = useState<HTMLButtonElement | null>(null);
    const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        setAnchorEl(event.currentTarget);
    };
    const handleClose = () => {
        setAnchorEl(null);
    };
    return (
        <div className={classes.account}>
            <Typography className={classes.accountName} variant="button" onClick={handleClick}>
                <NameBox name={props.child.name}/>
            </Typography>
            <MenuDropDown anchorEl={anchorEl} handleClose={handleClose} />
        </div>
    );
}
