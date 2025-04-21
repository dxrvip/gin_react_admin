import React, { useEffect } from 'react';
import {
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  ListSubheader,
  Stack,
  Typography
} from '@mui/material';
import { useChoicesContext, useInput } from 'react-admin';

const SelectProductField: React.FC<any> = (props) => {
  const { source, label } = props;
  const { allChoices } = useChoicesContext();
  const { field } = useInput({ source: source! });

  // 类型一致性处理
//   useEffect(() => {
//     if (typeof field.value === 'string') {
//       field.onChange(Number(field.value));
//     }
//   }, []);

  if (!allChoices) {
    return <div>Loading products...</div>;
  }

  const selectChildren: React.ReactNode[] = [];
  let isValidValue = false;

  allChoices.forEach((item: any) => {
    // 检查当前值是否有效
    const hasValidSku = item.secondHandSku?.some(
      (sku: any) => sku.id == field.value
    );
    if (hasValidSku) isValidValue = true;

    selectChildren.push(
      <ListSubheader key={item.id}>
        <Stack direction="row" spacing={2}>
          <Typography>{item.title}</Typography>
        </Stack>
      </ListSubheader>
    );

    item.secondHandSku?.forEach((sku: any) => {
      selectChildren.push(
        <MenuItem key={sku.id} value={Number(sku.id)}>
          <Stack direction="row" spacing={1}>
            <Typography>{sku.productsType}</Typography>
            <Typography>{sku.title}</Typography>
            <Typography>¥{sku.price}</Typography>
            <Typography>{sku.stock}</Typography>
          </Stack>
        </MenuItem>
      );
    });
  });
  const currentValue = field.value ?? '';
  return (
    <FormControl variant="filled" sx={{ m: 1, minWidth: 320 }} fullWidth>
      <InputLabel>选择商品</InputLabel>
      <Select
        value={field.value ?? ''}
        onChange={(e) => field.onChange(Number(e.target.value))}
        variant="filled"
      >
        {!isValidValue && field.value && (
          <MenuItem disabled value={field.value}>
            无效商品 (ID: {field.value})
          </MenuItem>
        )}
        {selectChildren}
      </Select>
    </FormControl>
  );
};

export default SelectProductField;