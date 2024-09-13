import { Container, Stack } from '@chakra-ui/react'
import TaskForm from './components/TaskForm'
import TaskList from './components/TaskList'

function App() {
   return (
    <Stack h="100-vh">
      <Container>
        <TaskForm/>
        <TaskList />
      </Container>
    </Stack>  
  ) 
}

export default App
