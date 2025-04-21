import {
  List,
  Datagrid,
  TextField,
  NumberField,
  useRecordContext,
  SelectField,
  DateField,
  useListContext,
} from 'react-admin';
import { Box, Dialog, DialogActions, DialogContent, DialogTitle, Button } from '@mui/material';
import { useEffect, useRef, useState } from 'react';
import { SECONDHAND_STATUS, FUNCTION_CHOICES, USAGE_DURATION_CHOICES, CONDITION_CHOICES, SecondHandProductCreate } from './SecondHandProductCreate';

export const SecondHandProductList = (props: any) => {
  const record = useRecordContext()
  const listRef = useRef<HTMLDivElement>(null);
  const [open, setOpen] = useState(false); // 确保对话框打开状态
  // const [setF] = useListContext()
  const handleClose = () => setOpen(false);
  const handleOpen = () => setOpen(true);
  useEffect(() => {
    if (listRef.current) {
      // 确保焦点不会留在设置了 aria-hidden 的元素内部
      listRef.current.tabIndex = -1;
      listRef.current.focus();
    }
  }, []);


  return (
    <Box ref={listRef}>
      <Button onClick={handleOpen} variant="text">打开</Button>

      <Dialog
        open={open}
        maxWidth="md" // 不设置最大宽度
        fullWidth={true} // 宽度占满
        onClose={handleClose}>
        <DialogTitle>商品sku</DialogTitle>
        <DialogContent>
          <List
            resource="secondHandSkus"
            filter={{ "product_id": record?.id }}
            sort={{ field: 'id', order: 'ASC' }}
            disableSyncWithLocation  // 新增：禁用与URL参数的同步
            filters={undefined}  // 新增：清除父级传递的过滤器
          >
            <Datagrid size="small" rowClick={false}>
              <TextField source="productsType" label="货号" />
              <DateField source="createAd" label="创建时间" />
              <SelectField
                source="condition"
                label="成色"
                choices={CONDITION_CHOICES}
              />
              <SelectField
                source="function"
                label="功能"
                choices={FUNCTION_CHOICES}
              />
              <SelectField
                source="usageDuration"
                label="使用时间"
                choices={USAGE_DURATION_CHOICES}
              />

              <NumberField source="price" options={{ style: 'currency', currency: 'CNY' }} label="代发价" />
              <NumberField source="cost" options={{ style: 'currency', currency: 'CNY' }} label="成本价" />
              <NumberField source="stock" label="库存" />
              <SelectField
                source="status"
                label="状态"
                choices={SECONDHAND_STATUS}
              />

            </Datagrid>
          </List>
        </DialogContent>
        <DialogActions>
          <SecondHandProductCreate />
          <Button onClick={handleClose}>关闭</Button>
        </DialogActions>
      </ Dialog>

    </Box>
  );
}


