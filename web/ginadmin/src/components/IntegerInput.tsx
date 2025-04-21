import { NumberInput, useInput } from 'react-admin';
import { TextField, InputAdornment } from '@mui/material';

// 自定义整数输入组件
const IntegerInput = (props: any) => {
  const {
    field,
    fieldState: { error }
  } = useInput({
    ...props,
    // 值转换逻辑（移除小数点）
    parse: (value) => value ? Math.floor(Number(value)) : null,
    // 格式化为整数
    format: (value) => value ? String(Math.floor(value)) : ''
  });

  return (
    <NumberInput
      {...props}
      InputProps={{
        // 禁止小数键盘（移动端优化）
        inputMode: 'numeric',
        pattern: '[0-9]*',
        // 添加后缀（可选）
        endAdornment: (
          <InputAdornment position="end">台</InputAdornment>
        )
      }}
      inputProps={{
        // 阻止小数点输入
        onKeyPress: (e) => {
          if (e.key === '.' || e.key === ',') {
            e.preventDefault();
          }
        },
        // 步进设置为1（键盘箭头控制）
        step: 1
      }}
      // 自定义 TextField 以显示验证错误
      TextFieldProps={{
        error: !!error,
        helperText: error?.message
      }}
      {...field}
    />
  );
};


export default IntegerInput