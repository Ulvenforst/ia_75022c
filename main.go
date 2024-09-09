package main

import (
  agentS "github.com/Ulvenforst/ia_750022c/agentesSimples"
)

func main() {
  testSimpleAgent()
}

func testSimpleAgent() {
  // En un agente simple a cada percepción se le asocia una acción. 

  // 1 es un obstáculo y 2 es el objetivo.
  envMatrix := [][]int{
    {2,0,0,0},
    {0,1,0,0},
    {0,1,0,0},
    {0,0,0,1},
  }
  enviroment := agentS.NewEnviroment(envMatrix)
  // El conocimiento se representa mediante una tabla.
  knowledge := map[[5]int]func(a *agentS.SimpleAgent){
    {1, 1, 1, 1, 0}: func(a *agentS.SimpleAgent) { a.MoveUp() },
    {1, 1, 0, 1, 0}: func(a *agentS.SimpleAgent) { a.MoveUp() },
    {1, 0, 1, 1, 0}: func(a *agentS.SimpleAgent) { a.MoveUp() },
    {1, 0, 0, 1, 0}: func(a *agentS.SimpleAgent) { a.MoveLeft() },
    {0, 1, 1, 1, 0}: func(a *agentS.SimpleAgent) { a.MoveLeft() },
    {0, 1, 0, 1, 0}: func(a *agentS.SimpleAgent) { a.MoveRight() },
    {0, 0, 1, 1, 0}: func(a *agentS.SimpleAgent) { a.MoveLeft() },
    {0, 0, 0, 1, 0}: func(a *agentS.SimpleAgent) { a.MoveLeft() },
    {1, 1, 1, 0, 0}: func(a *agentS.SimpleAgent) { a.MoveUp() },
    {1, 1, 0, 0, 0}: func(a *agentS.SimpleAgent) { a.MoveRight() },
    {1, 0, 1, 0, 0}: func(a *agentS.SimpleAgent) { a.MoveDown() },
    {1, 0, 0, 0, 0}: func(a *agentS.SimpleAgent) { a.MoveUp() },
    {0, 1, 1, 0, 0}: func(a *agentS.SimpleAgent) { a.MoveRight() },
    {0, 1, 0, 0, 0}: func(a *agentS.SimpleAgent) { a.MoveRight() },
    {0, 0, 1, 0, 0}: func(a *agentS.SimpleAgent) { a.MoveDown() },
  }
  agente := agentS.NewSimpleAgent(1,2, knowledge)

  // agente.GeneratePerception(*enviroment)
  // agente.LookForGoal(*enviroment)
  agente.VisualizePath(*enviroment)
}
