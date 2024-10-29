import React, { useCallback, ReactNode } from "react";
import useMediaQuery from '@mui/material/useMediaQuery';
import { useTheme } from '@mui/material/styles';
import {
    DialogContent,
    DialogTitle,
    Dialog,
    DialogContentText,
    DialogActions,
} from "@mui/material";

export interface DialogWindowProps {
    open: boolean;
    onClose: (b: boolean) => void;
    children: ReactNode;
    contentText?: string;
    title?: string;
    dialogActions?: ReactNode;
}

const Dialogwindow = (props: DialogWindowProps) => {
    const theme = useTheme();
    const fullScreen = useMediaQuery(theme.breakpoints.down('sm'));
    const {
        open,
        onClose,
        children,
        contentText,
        title,
        dialogActions,
    } = props;

    const handleClose = useCallback(() => {
        onClose(false);
    }, [onClose]);


    return (
        <Dialog onClose={handleClose} open={open} fullScreen={fullScreen}>
            {title && <DialogTitle>{title}</DialogTitle>}
            <DialogContent>
                {contentText && <DialogContentText id="alert-dialog-slide-description">{contentText}</DialogContentText>}
                {children}  {/* 修正为 children */}
            </DialogContent>
            {dialogActions && <DialogActions>
                {dialogActions}
            </DialogActions>}
        </Dialog>
    );
};

export default Dialogwindow;