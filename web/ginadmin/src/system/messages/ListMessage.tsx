import {
    List,
    Datagrid,
    TextField,
    DateField,
    
}from 'react-admin'


function ListMessage() {
    return (  
        <List>
            <Datagrid>
                <TextField source="id"/>
                <TextField source="title" label="标题"/>
                <TextField source="userName" label="创建人"/>
                <DateField source="createTime" label="创建时间" />
                <TextField source="status" label="状态" />
            </Datagrid>
        </List>
    );
}

export default ListMessage;