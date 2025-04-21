import * as React from 'react';
import { useRecordContext, useInput, FieldProps, useGetList } from 'react-admin';
import { SimpleTreeView } from '@mui/x-tree-view/SimpleTreeView';
import { TreeItem } from '@mui/x-tree-view/TreeItem';
import { Checkbox, FormControl, FormControlLabel, Box } from '@mui/material';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';

// 类型定义
export interface TreeNode {
  id: string;
  name: string;
  children?: TreeNode[];
}

interface TreeInputProps extends FieldProps {
  source: string;
  multiple?: boolean;
  defaultExpanded?: string[];
}

const TreeInput: React.FC<TreeInputProps> = ({ 
  source, 
  multiple = true,
  defaultExpanded = []
}) => {
  const record = useRecordContext();
  const { field } = useInput({ source });
  const [selectedIds, setSelectedIds] = React.useState<string[]>(
    record?.[source] || []
  );

  // 获取分类数据
  const { data: categories } = useGetList<TreeNode>('categories', {
    pagination: { page: 1, perPage: 1000 },
    filter: { parent_id_is_null: true }
  });

  // 递归渲染树节点
  const renderTree = (nodes: TreeNode) => (
    <TreeItem
      key={nodes.id}
      itemId={nodes.id}
      label={
        <FormControlLabel
          control={
            <Checkbox
              checked={selectedIds.includes(nodes.id)}
              onClick={(e) => e.stopPropagation()}
              onChange={(e) => handleSelect(e, nodes.id)}
            />
          }
          label={nodes.name}
          sx={{ userSelect: 'none' }}
        />
      }
    >
      {nodes.children?.map((node) => renderTree(node))}
    </TreeItem>
  );

  // 处理选择
  const handleSelect = (
    event: React.ChangeEvent<HTMLInputElement>,
    nodeId: string
  ) => {
    const newSelectedIds = multiple
      ? selectedIds.includes(nodeId)
        ? selectedIds.filter(id => id !== nodeId)
        : [...selectedIds, nodeId]
      : [nodeId];

    setSelectedIds(newSelectedIds);
    field.onChange(newSelectedIds);
  };

  return (
    <FormControl component="fieldset" fullWidth>
      <Box sx={{ maxHeight: 400, overflow: 'auto', p: 1 }}>
        <SimpleTreeView
          aria-label="category-tree"
          defaultCollapseIcon={<ExpandMoreIcon />}
          defaultExpandIcon={<ChevronRightIcon />}
          defaultExpandedItems={defaultExpanded}
          multiSelect={multiple}
          selectedItems={selectedIds}
        >
          {categories?.map(node => renderTree(node))}
        </SimpleTreeView>
      </Box>
    </FormControl>
  );
};

export default TreeInput;