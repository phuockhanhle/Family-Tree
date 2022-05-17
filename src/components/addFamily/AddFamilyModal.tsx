import { Button, Dialog, DialogActions, DialogContent, DialogTitle, FormControl, Input, InputLabel, List, ListItem } from "@mui/material";
import { useState } from "react";


interface AddFamilyModalProps {
    isOpen: boolean;
    onClose: () => void;
    onSubmit: (familyMemberName: string) => void;
}

export const AddFamilyModal: React.FC<AddFamilyModalProps> = ({
    isOpen,
    onClose,
    onSubmit
}) => {
    const [name, setName] = useState('');
    return <Dialog open={isOpen} onClose={onClose}>
        <DialogTitle>Add family member</DialogTitle>
        <DialogContent>
            <FormControl>
                <InputLabel>Family member</InputLabel>
                <Input value={name} onChange={(event) => setName(event.target.value)} />
            </FormControl>
        </DialogContent>
        <DialogActions>
            <Button sx={{
                color: 'blue'
            }}
                disabled={!name}
                onClick={() => onSubmit(name)}>
                Add
            </Button>
        </DialogActions>
    </Dialog>
}