import { MenuItem, Popover } from "@material-ui/core"
import { useNavigate } from "react-router-dom";
import { useStyles } from "../../utils/defaultStyles"
import { openNewTabWithLink } from "../../utils/helperFunc";

export const MenuDropDown = (props: any) => {
    const classes = useStyles();
    const navigate = useNavigate();

    const open = Boolean(props.anchorEl);
    const id = open ? 'simple-popover' : undefined;

    function goToPage(link: string) {
        navigate(link);
    };

    function doSomething() {
        return
    }

    return (
        <Popover className={classes.popOver}
            id={id}
            open={open}
            anchorEl={props.anchorEl}
            onClose={props.handleClose}
            anchorOrigin={{
                vertical: 'bottom',
                horizontal: 'right',
            }}
            transformOrigin={{
                vertical: 'top',
                horizontal: 'right',
            }}>
            {/* {
                props.parent.map(
                    (role: any) => {
                        return (<MenuItem key={role} className={classes.popOverItem}
                            onClick={() => doSomething()}>
                            <img src="../static/pin.png" width="30" height="30" alt="pin" />
                            <p style={{ marginLeft: 10 }}> {role.name} </p>
                        </MenuItem>)
                    })
            } */}
            <MenuItem className={classes.popOverItem}
                onClick={() => doSomething()}>
                <p style={{ marginLeft: 10 }}>Parent</p>
            </MenuItem>
            <MenuItem className={classes.popOverItem}
                onClick={() => doSomething()}>

                <p style={{ marginLeft: 10 }}>Child</p>
            </MenuItem>
            <MenuItem className={classes.popOverItem}
                onClick={() => doSomething()}>
                <p style={{ marginLeft: 10 }}>Gender</p>
            </MenuItem>
        </Popover>)
}

