
import {useRecordContext} from "react-admin"
import Chip from '@mui/material/Chip';
export default function StatusField(prop: any) {
    const record = useRecordContext()

    if(!record) return null 

    function C() {
        switch (record?.status) {
            case "发布":
                return "success"
            case "草稿":
                return "primary"
            case "下架":
                return "warning"
            default:
                return "info"
        }
    }
    return (
        <Chip label={record.status} variant="outlined" color={(C() as any)}  />
    )
}