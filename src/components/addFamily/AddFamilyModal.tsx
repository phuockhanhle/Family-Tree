import { Button, Dialog, DialogActions, DialogContent, DialogTitle, FormControl, Input, InputLabel, List, ListItem, ListItemButton, ListItemIcon, ListItemText } from "@mui/material";
import { useState } from "react";
import AddIcon from '@mui/icons-material/Add';

interface AddFamilyModalProps {
    isOpen: boolean;
    onClose: () => void;
    onSubmit: (familyMemberName: string, identity: number) => void;
    onClickAddFamily: (familyTree: any) => void;
    familyTree: any;
}

export const AddFamilyModal: React.FC<AddFamilyModalProps> = ({
    isOpen,
    onClose,
    onSubmit,
    onClickAddFamily,
    familyTree
}) => {
    const [name, setName] = useState('');
    const [identity, setIdentity] = useState(0);



    return <Dialog open={isOpen} onClose={onClose}>
        <DialogTitle>Add family member</DialogTitle>
        <DialogContent>
            <List>
                <ListItem>
                    <FormControl>
                        <InputLabel>Family member</InputLabel>
                        <Input value={name} onChange={(event) => setName(event.target.value)} />
                    </FormControl>
                </ListItem>
                <ListItem>
                    <FormControl>
                        <InputLabel>Id</InputLabel>
                        <Input value={identity} onChange={(event) => setIdentity(Number(event.target.value))} />
                    </FormControl>
                </ListItem>
                <DialogActions>
                    <Button variant='contained'
                        disabled={!name}
                        onClick={() => {
                            onSubmit(name, identity);
                            onClose();
                        }}>
                        Add
                    </Button>
                </DialogActions>
                <ListItemButton onClick={() => {
                    onClickAddFamily(familyTree);
                    onClose();
                }}>
                    <ListItemIcon>
                        <AddIcon />
                    </ListItemIcon>
                    <ListItemText>Add Family Tree</ListItemText>
                </ListItemButton>
            </List>
        </DialogContent>

    </Dialog>
}