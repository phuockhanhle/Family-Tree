import Avatar from '@mui/material/Avatar';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemAvatar from '@mui/material/ListItemAvatar';
import ListItemText from '@mui/material/ListItemText';
import DialogTitle from '@mui/material/DialogTitle';
import Dialog from '@mui/material/Dialog';
import PersonIcon from '@mui/icons-material/Person';
import AddIcon from '@mui/icons-material/Add';
import { blue } from '@mui/material/colors';

interface SimpleDialogProps {
    open: boolean;
    selectedValue: string;
    onClose: (value: string) => void;
}
const options = [
    'Add person',
    'Remove person',
    'Change Family'
]

export const SimpleDialog = (props: SimpleDialogProps) => {
    const { onClose, selectedValue, open } = props;

    const handleClose = () => {
        onClose(selectedValue);
    };

    const handleListItemClick = (value: string) => {
        onClose(value);
    };

    return (
        <Dialog onClose={handleClose} open={open}>
            <DialogTitle>Adjust FamilyTree</DialogTitle>
            <List sx={{ pt: 0 }}>
                {options.map((option) => (
                    <ListItem button onClick={() => handleListItemClick(option)} key={option}>
                        <ListItemAvatar>
                            <Avatar sx={{ bgcolor: blue[100], color: blue[600] }}>
                                <AddIcon />
                            </Avatar>
                        </ListItemAvatar>
                        <ListItemText primary={option} />
                    </ListItem>
                ))}
            </List>
        </Dialog>
    );
}