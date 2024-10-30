// theme.ts
import { createTheme } from '@mui/material/styles';
import type {} from '@mui/lab/themeAugmentation';


const theme = createTheme({
  components: {
    MuiTimeline: {
      styleOverrides: {
        root: {
          backgroundColor: 'red', // 修改根样式
        },
      },
    },
  },
});

export default theme;