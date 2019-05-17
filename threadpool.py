#coding:utf8
from threading import Thread
import threading
import time
from Queue import Queue

class ThreadPoolManger():
	def __init__(self, size):
		self.poolsize = size
		self.pool = Queue()
		self.__initThreadPool(self.poolsize)

	def __initThreadPool(self, size):
		for i in xrange(0, size):
			t = ThreadIns(self.pool)
			t.start()

	def add_job(self, func):
		self.pool.put(func)


class ThreadIns(Thread):
	def __init__(self, pool):
		Thread.__init__(self)
		self.workPool = pool

	def run(self):
		while True:
			target = self.workPool.get()
			target()
			self.workPool.task_done()

def demo_func():
	print 'thread %s is running ' % threading.current_thread().name

pool = ThreadPoolManger(5)

for i in xrange(0, 10):
	pool.add_job(demo_func)

time.sleep(120)