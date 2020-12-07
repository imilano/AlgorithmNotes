package _6_greedy

import (
	"sort"
)

type task struct {
	deadline int
	punish int
}

type Tasks []task

func (t Tasks)Less(i,j int) bool {
	return t[i].punish > t[j].punish
}

func (t Tasks) Len() int {
	return len(t)
}

func (t Tasks)Swap(i,j int) {
	t[i],t[j] = t[j],t[i]
}

func schedule(deadline []int,punish []int) (int,[]task) {
	var res []task
	var totalPunish int
	var unscheduled []task
	hasScheduled := make(map[int]*task)

	var tasks Tasks
	for i := range deadline {
		t := task{
			deadline: deadline[i],
			punish:   punish[i],
		}

		tasks = append(tasks,t)
	}

	// 按照punish从大到小排序
	sort.Sort(tasks)

	// 从左到右遍历每个任务，将其安排在离截止时间点最近（包括截止时间点处）的未安排任务的时间点完成，若无法找到这个时间点，
	// 则这个任务无法按时完成，将其误时惩罚加到总误时惩罚中。
	for i := 0; i <len(tasks);i++ {
		visited := false
		for j := tasks[i].deadline; j > 0;j-- {
			if hasScheduled[j] == nil {
				hasScheduled[j] = &tasks[i]
				visited = true
				break
			}
		}

		if !visited {
			totalPunish += tasks[i].punish
			unscheduled = append(unscheduled,tasks[i])
		}
	}

	for _,v := range hasScheduled {
		res = append(res,*v)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].deadline < res[j].deadline
	})
	for _,v := range unscheduled {
		res = append(res,v)
	}
	return totalPunish,res
}