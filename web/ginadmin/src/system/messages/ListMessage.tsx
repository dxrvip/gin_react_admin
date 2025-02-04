import {
    List,
    Datagrid,
    TextField,
    NumberField,
    
}from 'react-admin'


function ListMessage() {
    return (  
        <List>
            <Datagrid>
                <TextField source="id"/>
                <TextField source="message" label="创建人"/>
                <TextField source="createAd" label="创建时间" />
                <TextField source="updatedAd" label="更新时间"/>
                <TextField source="status" label="状态" />
            </Datagrid>
        </List>
    );
}

export default ListMessage;