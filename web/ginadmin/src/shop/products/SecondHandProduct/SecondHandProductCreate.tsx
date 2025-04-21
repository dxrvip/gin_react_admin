import { useState  } from "react";
import {
    Create,

    useCreateSuggestionContext,

    useRecordContext,

} from "react-admin";
import { Box, Button, Dialog, DialogContent, DialogTitle } from "@mui/material";
import { FormSeconHand } from "./FormSeconHand";


export const FUNCTION_CHOICES = [
    { id: 'perfect', name: '功能完好无维修' },
    { id: 'repaired', name: '维修过' },
    { id: 'usable', name: '可以正常使用' },
    { id: 'unusable', name: '无法正常使用' },
];

export const USAGE_DURATION_CHOICES = [
    { id: 'unused', name: '未使用' },
    { id: 'half_year', name: '6个月内' },
    { id: 'one_year', name: '6个月-1年' },
    { id: 'three_years', name: '1-2年' },
];
// 分离 choices 对象
export const CONDITION_CHOICES = [
    { id: 'new', name: '全新' },
    { id: 'like_new', name: '几乎全新' },
    { id: 'light', name: '轻微痕迹' },
    { id: 'obvious', name: '明显痕迹' },
    { id: 'serious', name: '严重痕迹' },
    { id: 'damaged', name: '破损' },
];
// 添加是否再保修
const IS_REPAIR_CHOICES = [
    { id: 'repair', name: '再保修' },
    { id: 'not_repair', name: '不保修' },
];

export const SECONDHAND_STATUS = [
    { id: 'active', name: '发布' },
    { id: 'disabled', name: '草稿' },
    { id: 'pulled', name: '下架' },
    // { id: 'deleted', name: '删除'},

    // 增加卖完
    // { id: 'sold_out', name: '售罄' },
]
// 添加二手商品表单
export const SecondHandProductCreate = () => {
    const record = useRecordContext()
    const [open, setOpen] = useState(false);
    const handleOpen = () => setOpen(true);
    const handleClose = () => setOpen(false);

    if (!record) return <Box>无数据。。。</Box>
    return (
        <>
            <Button onClick={handleOpen}>添加</Button>
            <Dialog open={open} onClose={handleClose} fullWidth maxWidth="md">
                <DialogTitle>添加二手商品</DialogTitle>
                <DialogContent>
                    <Box sx={{ width: '100%' }}>

                        <Create
                            redirect={false}>
                                <FormSeconHand record={{...record, handleClose}} />
                            
                        </Create>
                    </Box>
                </DialogContent>

            </Dialog>
        </>
    )
}

